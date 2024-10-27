// @generated by protoc-gen-connect-es v1.6.1
// @generated from file version/v1/version.proto (package version, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetCurrentVersionRequest, GetCurrentVersionResponse, GetVersionsRequest, GetVersionsResponse } from "./version_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service version.VersionService
 */
export const VersionService = {
  typeName: "version.VersionService",
  methods: {
    /**
     * @generated from rpc version.VersionService.GetVersions
     */
    getVersions: {
      name: "GetVersions",
      I: GetVersionsRequest,
      O: GetVersionsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc version.VersionService.GetCurrentVersion
     */
    getCurrentVersion: {
      name: "GetCurrentVersion",
      I: GetCurrentVersionRequest,
      O: GetCurrentVersionResponse,
      kind: MethodKind.Unary,
    },
  }
};
