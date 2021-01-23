import gql from 'graphql-tag';

export const ALL_POST_QUERY = gql`
query GetAllPosts($orderBy: OrderBy, $first: Int!, $skip: Int!) {
  allPosts(orderBy: $orderBy, first: $first, skip: $skip) {
    id
    title
  }
}
`;