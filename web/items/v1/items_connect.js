// @generated by protoc-gen-connect-es v1.6.1
// @generated from file items/v1/items.proto (package items, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetBaseItemsRequest, GetBaseItemsResponse, GetItemImageByIdRequest, GetItemImageByIdResponse, GetItemInformationByIdRequest, GetItemInformationByIdResponse } from "./items_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service items.ItemsService
 */
export const ItemsService = {
  typeName: "items.ItemsService",
  methods: {
    /**
     * @generated from rpc items.ItemsService.GetItemImageById
     */
    getItemImageById: {
      name: "GetItemImageById",
      I: GetItemImageByIdRequest,
      O: GetItemImageByIdResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc items.ItemsService.GetItemInformationById
     */
    getItemInformationById: {
      name: "GetItemInformationById",
      I: GetItemInformationByIdRequest,
      O: GetItemInformationByIdResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc items.ItemsService.GetBaseItems
     */
    getBaseItems: {
      name: "GetBaseItems",
      I: GetBaseItemsRequest,
      O: GetBaseItemsResponse,
      kind: MethodKind.Unary,
    },
  }
};
