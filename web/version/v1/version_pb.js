// @generated by protoc-gen-es v1.10.0
// @generated from file version/v1/version.proto (package version, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message version.GetVersionsRequest
 */
export const GetVersionsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "version.GetVersionsRequest",
  [],
);

/**
 * @generated from message version.GetVersionsResponse
 */
export const GetVersionsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "version.GetVersionsResponse",
  () => [
    { no: 1, name: "versions", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message version.GetCurrentVersionRequest
 */
export const GetCurrentVersionRequest = /*@__PURE__*/ proto3.makeMessageType(
  "version.GetCurrentVersionRequest",
  [],
);

/**
 * @generated from message version.GetCurrentVersionResponse
 */
export const GetCurrentVersionResponse = /*@__PURE__*/ proto3.makeMessageType(
  "version.GetCurrentVersionResponse",
  () => [
    { no: 1, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

