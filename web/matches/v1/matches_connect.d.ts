// @generated by protoc-gen-connect-es v1.6.1
// @generated from file matches/v1/matches.proto (package matches, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetMatchByMatchIdRequest, GetMatchByMatchIdResponse, GetMatchIdsByPuuidRequest, GetMatchIdsByPuuidResponse, GetMatchTimelineRequest, GetMatchTimelineResponse } from "./matches_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service matches.MatchesService
 */
export declare const MatchesService: {
  readonly typeName: "matches.MatchesService",
  readonly methods: {
    /**
     * @generated from rpc matches.MatchesService.GetMatchTimeline
     */
    readonly getMatchTimeline: {
      readonly name: "GetMatchTimeline",
      readonly I: typeof GetMatchTimelineRequest,
      readonly O: typeof GetMatchTimelineResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc matches.MatchesService.GetMatchByMatchId
     */
    readonly getMatchByMatchId: {
      readonly name: "GetMatchByMatchId",
      readonly I: typeof GetMatchByMatchIdRequest,
      readonly O: typeof GetMatchByMatchIdResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc matches.MatchesService.GetMatchIdsByPuuid
     */
    readonly getMatchIdsByPuuid: {
      readonly name: "GetMatchIdsByPuuid",
      readonly I: typeof GetMatchIdsByPuuidRequest,
      readonly O: typeof GetMatchIdsByPuuidResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

