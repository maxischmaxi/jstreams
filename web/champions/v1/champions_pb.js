// @generated by protoc-gen-es v1.10.0
// @generated from file champions/v1/champions.proto (package champions, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message champions.GetChampByChampIdRequest
 */
export const GetChampByChampIdRequest = /*@__PURE__*/ proto3.makeMessageType(
  "champions.GetChampByChampIdRequest",
  () => [
    { no: 1, name: "patchVersion", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "champId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message champions.GetChampByChampIdResponse
 */
export const GetChampByChampIdResponse = /*@__PURE__*/ proto3.makeMessageType(
  "champions.GetChampByChampIdResponse",
  () => [
    { no: 1, name: "champion", kind: "message", T: Champion },
  ],
);

/**
 * @generated from message champions.GetChampionsRequest
 */
export const GetChampionsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "champions.GetChampionsRequest",
  () => [
    { no: 1, name: "patchVersion", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message champions.GetChampionsResponse
 */
export const GetChampionsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "champions.GetChampionsResponse",
  () => [
    { no: 1, name: "type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "format", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "data", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Champion} },
  ],
);

/**
 * @generated from message champions.ChampionInfo
 */
export const ChampionInfo = /*@__PURE__*/ proto3.makeMessageType(
  "champions.ChampionInfo",
  () => [
    { no: 1, name: "attack", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "defense", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "magic", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "difficulty", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message champions.Champion
 */
export const Champion = /*@__PURE__*/ proto3.makeMessageType(
  "champions.Champion",
  () => [
    { no: 1, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "blurb", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "info", kind: "message", T: ChampionInfo },
    { no: 8, name: "image", kind: "message", T: Image },
    { no: 9, name: "tags", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 10, name: "partype", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "stats", kind: "message", T: ChampionStats },
  ],
);

/**
 * @generated from message champions.ChampionStats
 */
export const ChampionStats = /*@__PURE__*/ proto3.makeMessageType(
  "champions.ChampionStats",
  () => [
    { no: 1, name: "hp", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "hpperlevel", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "mp", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "mpperlevel", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 5, name: "movespeed", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "armor", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 7, name: "armorperlevel", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 8, name: "spellblock", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 9, name: "spellblockperlevel", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 10, name: "attackrange", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 11, name: "hpregen", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 12, name: "hpregenperlevel", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 13, name: "mpregen", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 14, name: "mpregenperlevel", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 15, name: "crit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 16, name: "critperlevel", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 17, name: "attackdamage", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 18, name: "attackdamageperlevel", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 19, name: "attackspeed", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 20, name: "attackspeedperlevel", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ],
);

/**
 * @generated from message champions.Image
 */
export const Image = /*@__PURE__*/ proto3.makeMessageType(
  "champions.Image",
  () => [
    { no: 1, name: "full", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "sprite", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "group", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "x", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "y", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "w", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 7, name: "h", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

