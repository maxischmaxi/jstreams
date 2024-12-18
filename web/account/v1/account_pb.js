// @generated by protoc-gen-es v1.10.0
// @generated from file account/v1/account.proto (package account, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum account.RegionalRoutingValues
 */
export const RegionalRoutingValues = /*@__PURE__*/ proto3.makeEnum(
  "account.RegionalRoutingValues",
  [
    {no: 0, name: "EUROPE"},
    {no: 1, name: "AMERICAS"},
    {no: 2, name: "ASIA"},
    {no: 3, name: "ESPORTS"},
  ],
);

/**
 * @generated from message account.GetAccountByGamenameAndTaglineRequest
 */
export const GetAccountByGamenameAndTaglineRequest = /*@__PURE__*/ proto3.makeMessageType(
  "account.GetAccountByGamenameAndTaglineRequest",
  () => [
    { no: 1, name: "gamename", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "tagline", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "region", kind: "enum", T: proto3.getEnumType(RegionalRoutingValues) },
  ],
);

/**
 * @generated from message account.Account
 */
export const Account = /*@__PURE__*/ proto3.makeMessageType(
  "account.Account",
  () => [
    { no: 1, name: "puuid", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "gameName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "tagLine", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message account.GetAccountByGamenameAndTaglineResponse
 */
export const GetAccountByGamenameAndTaglineResponse = /*@__PURE__*/ proto3.makeMessageType(
  "account.GetAccountByGamenameAndTaglineResponse",
  () => [
    { no: 1, name: "account", kind: "message", T: Account },
  ],
);

/**
 * @generated from message account.GetAccountProfileIconRequest
 */
export const GetAccountProfileIconRequest = /*@__PURE__*/ proto3.makeMessageType(
  "account.GetAccountProfileIconRequest",
  () => [
    { no: 1, name: "profileIconId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "patchVersion", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message account.GetAccountProfileIconResponse
 */
export const GetAccountProfileIconResponse = /*@__PURE__*/ proto3.makeMessageType(
  "account.GetAccountProfileIconResponse",
  () => [
    { no: 1, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message account.GetAccountByPuuidRequest
 */
export const GetAccountByPuuidRequest = /*@__PURE__*/ proto3.makeMessageType(
  "account.GetAccountByPuuidRequest",
  () => [
    { no: 1, name: "puuid", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "region", kind: "enum", T: proto3.getEnumType(RegionalRoutingValues) },
  ],
);

/**
 * @generated from message account.GetAccountByPuuidResponse
 */
export const GetAccountByPuuidResponse = /*@__PURE__*/ proto3.makeMessageType(
  "account.GetAccountByPuuidResponse",
  () => [
    { no: 1, name: "account", kind: "message", T: Account },
  ],
);

