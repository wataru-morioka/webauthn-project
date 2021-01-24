/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { OrderBy } from "./../graphql-types";

// ====================================================
// GraphQL query operation: GetAllPosts
// ====================================================

export interface GetAllPosts_allPosts {
  __typename: "Post";
  id: string;
  title: string;
}

export interface GetAllPosts {
  allPosts: GetAllPosts_allPosts[];
}

export interface GetAllPostsVariables {
  orderBy?: OrderBy | null;
  first: number;
  skip: number;
}
