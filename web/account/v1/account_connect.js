// @generated by protoc-gen-connect-es v1.6.1
// @generated from file account/v1/account.proto (package account, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetAccountByGamenameAndTaglineRequest, GetAccountByGamenameAndTaglineResponse, GetAccountByPuuidRequest, GetAccountByPuuidResponse, GetAccountProfileIconRequest, GetAccountProfileIconResponse } from "./account_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service account.AccountService
 */
export const AccountService = {
  typeName: "account.AccountService",
  methods: {
    /**
     * @generated from rpc account.AccountService.GetAccountByGamenameAndTagline
     */
    getAccountByGamenameAndTagline: {
      name: "GetAccountByGamenameAndTagline",
      I: GetAccountByGamenameAndTaglineRequest,
      O: GetAccountByGamenameAndTaglineResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc account.AccountService.GetAccountProfileIcon
     */
    getAccountProfileIcon: {
      name: "GetAccountProfileIcon",
      I: GetAccountProfileIconRequest,
      O: GetAccountProfileIconResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc account.AccountService.GetAccountByPuuid
     */
    getAccountByPuuid: {
      name: "GetAccountByPuuid",
      I: GetAccountByPuuidRequest,
      O: GetAccountByPuuidResponse,
      kind: MethodKind.Unary,
    },
  }
};

