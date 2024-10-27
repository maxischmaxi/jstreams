// @generated by protoc-gen-connect-es v1.6.1
// @generated from file masteries/v1/masteries.proto (package masteries, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetChampionMasteriesByChampionRequeset, GetChampionMasteriesByChampionResponse, GetChampionMasteriesRequeset, GetChampionMasteriesResponse } from "./masteries_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service masteries.MasteriesService
 */
export const MasteriesService = {
  typeName: "masteries.MasteriesService",
  methods: {
    /**
     * @generated from rpc masteries.MasteriesService.GetChampionMasteriesByPuuid
     */
    getChampionMasteriesByPuuid: {
      name: "GetChampionMasteriesByPuuid",
      I: GetChampionMasteriesRequeset,
      O: GetChampionMasteriesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc masteries.MasteriesService.GetChampionMasteriesByPuuidByChampion
     */
    getChampionMasteriesByPuuidByChampion: {
      name: "GetChampionMasteriesByPuuidByChampion",
      I: GetChampionMasteriesByChampionRequeset,
      O: GetChampionMasteriesByChampionResponse,
      kind: MethodKind.Unary,
    },
  }
};

