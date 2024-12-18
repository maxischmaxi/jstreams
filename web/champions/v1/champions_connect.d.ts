// @generated by protoc-gen-connect-es v1.6.1
// @generated from file champions/v1/champions.proto (package champions, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetChampByChampIdRequest, GetChampByChampIdResponse, GetChampionsRequest, GetChampionsResponse } from "./champions_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service champions.ChampionsService
 */
export declare const ChampionsService: {
  readonly typeName: "champions.ChampionsService",
  readonly methods: {
    /**
     * @generated from rpc champions.ChampionsService.GetChampions
     */
    readonly getChampions: {
      readonly name: "GetChampions",
      readonly I: typeof GetChampionsRequest,
      readonly O: typeof GetChampionsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc champions.ChampionsService.GetChampByChampId
     */
    readonly getChampByChampId: {
      readonly name: "GetChampByChampId",
      readonly I: typeof GetChampByChampIdRequest,
      readonly O: typeof GetChampByChampIdResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

