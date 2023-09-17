import { useQuery } from '@tanstack/react-query';
import { useAtom } from 'jotai';

import { apiClient } from '@/apiClient';
import { userAtom } from '@/atom';

import { HeaderWrapper, LeftWrapper, RightWrapper } from './style';

export function AppHeader() {
  const [user] = useAtom(userAtom);

  const { data } = useQuery({
    queryKey: ['user'],
    enabled: !!user.accessToken,
    queryFn: async () => {
      console.log(`Bearer ${user.accessToken}`);
      const response = await apiClient.get<{
        success: boolean;
        company: {
          name: string;
          description: string;
          domain: string;
        };
      }>('/api/company', {
        headers: {
          Authorization: `Bearer ${user.accessToken}`,
        },
      });
      return response.data;
    },
  });

  return (
    <HeaderWrapper>
      <LeftWrapper>{data?.company.name}</LeftWrapper>
      <RightWrapper
        onContextMenu={(e) => {
          e.preventDefault();
          console.log(e.clientX, e.clientY);
        }}>
        Profile
      </RightWrapper>
    </HeaderWrapper>
  );
}
