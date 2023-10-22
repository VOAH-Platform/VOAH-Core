import { apiClient } from '@/apiClient';

export async function getProfileById(id: string, accessToken: string) {
  const res = apiClient.get<{
    user: {
      'user-id': string;
      email: string;
      username: string;
      displayname: string;
      position: string;
      description: string;
      'team-id': string;
      roles: unknown;
      projects: unknown[];
      'created-at': string;
    };
  }>(`/api/profile?user-id=${id}`, {
    headers: {
      Authorization: `Bearer ${accessToken}`,
    },
  });
  return res;
}
