/*
 * Copyright 2018 The Feast Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package feast.ingestion.model;

import com.google.common.base.Preconditions;
import feast.ingestion.service.SpecService;
import feast.specs.EntitySpecProto.EntitySpec;
import feast.specs.FeatureSpecProto.FeatureSpec;
import feast.specs.ImportSpecProto.Field;
import feast.specs.ImportSpecProto.ImportSpec;
import feast.specs.StorageSpecProto.StorageSpec;
import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import lombok.Builder;
import lombok.Getter;

@Builder
@Getter
public class Specs implements Serializable {
  private String jobName;
  private ImportSpec importSpec;
  private Map<String, EntitySpec> entitySpecs;
  private Map<String, FeatureSpec> featureSpecs;
  private Map<String, StorageSpec> storageSpecs;
  private transient SpecService specService;
  private RuntimeException error;

  public static Specs of(String jobName, ImportSpec importSpec, SpecService specService) {
    try {
      Specs.SpecsBuilder specsBuilder = Specs.builder().jobName(jobName).importSpec(importSpec);

      List<Field> fields = importSpec.getSchema().getFieldsList();
      List<String> featureIds = new ArrayList<>();
      for (Field field : fields) {
        if (!field.getFeatureId().isEmpty()) {
          featureIds.add(field.getFeatureId());
        }
      }
      specsBuilder.featureSpecs(specService.getFeatureSpecs(featureIds));

      List<String> entityNames = importSpec.getEntitiesList();
      for (FeatureSpec featureSpec : specsBuilder.featureSpecs.values()) {
        Preconditions.checkArgument(
            entityNames.contains(featureSpec.getEntity()),
            "Feature has entity not listed in import spec featureSpec=" + featureSpec.toString());
      }
      specsBuilder.entitySpecs(specService.getEntitySpecs(entityNames));

      specsBuilder.storageSpecs(specService.getAllStorageSpecs());

      return specsBuilder.build();
    } catch (RuntimeException e) {
      return Specs.builder().error(e).build();
    }
  }

  public void validate() {
    if (error != null) {
      throw error;
    }

    // Sanity checks that our maps are built correctly
    for (Entry<String, FeatureSpec> entry : featureSpecs.entrySet()) {
      Preconditions.checkArgument(entry.getKey().equals(entry.getValue().getId()));
    }
    for (Entry<String, EntitySpec> entry : entitySpecs.entrySet()) {
      Preconditions.checkArgument(entry.getKey().equals(entry.getValue().getName()));
    }
    for (Entry<String, StorageSpec> entry : storageSpecs.entrySet()) {
      Preconditions.checkArgument(entry.getKey().equals(entry.getValue().getId()));
    }

    for (FeatureSpec featureSpec : featureSpecs.values()) {
      // Check that feature has a matching entity
      Preconditions.checkArgument(
          entitySpecs.containsKey(featureSpec.getEntity()),
          String.format(
              "Feature %s references unknown entity %s",
              featureSpec.getId(), featureSpec.getEntity()));
      // Check that feature has a matching serving store
      Preconditions.checkArgument(
          storageSpecs.containsKey(featureSpec.getDataStores().getServing().getId()),
          String.format(
              "Feature %s references unknown serving store %s",
              featureSpec.getId(), featureSpec.getDataStores().getServing().getId()));
      // Check that feature has a matching warehouse store
      Preconditions.checkArgument(
          storageSpecs.containsKey(featureSpec.getDataStores().getWarehouse().getId()),
          String.format(
              "Feature %s references unknown warehouse store %s",
              featureSpec.getId(), featureSpec.getDataStores().getWarehouse().getId()));
    }
  }

  public EntitySpec getEntitySpec(String entityName) {
    Preconditions.checkArgument(
        entitySpecs.containsKey(entityName),
        String.format("Unknown entity %s, spec was not initialized", entityName));
    return entitySpecs.get(entityName);
  }

  public FeatureSpec getFeatureSpec(String featureId) {
    Preconditions.checkArgument(
        featureSpecs.containsKey(featureId),
        String.format("Unknown feature %s, spec was not initialized", featureId));
    return featureSpecs.get(featureId);
  }

  public List<FeatureSpec> getFeatureSpecByServingStoreId(String storeId) {
    List<FeatureSpec> out = new ArrayList<>();
    for (FeatureSpec featureSpec : featureSpecs.values()) {
      if (featureSpec.getDataStores().getServing().getId().equals(storeId)) {
        out.add(featureSpec);
      }
    }
    return out;
  }

  public StorageSpec getStorageSpec(String storeId) {
    Preconditions.checkArgument(
        storageSpecs.containsKey(storeId),
        String.format("Unknown store %s, spec was not initialized", storeId));
    return storageSpecs.get(storeId);
  }
}
