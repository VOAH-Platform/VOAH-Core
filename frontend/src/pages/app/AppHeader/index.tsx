import { useQuery } from '@tanstack/react-query';
import { useAtom } from 'jotai';
import { useEffect } from 'react';

import { API_HOST, apiClient } from '@/apiClient';
import { userAtom } from '@/atom';

import { MD5 } from './md5';
import {
  CompanyName,
  HeaderWrapper,
  ProfileWrapper,
  LeftWrapper,
  RightWrapper,
  StatusMargin,
  Status,
  ImageWrapper,
  CompanyWrapper,
  CompanyImage,
} from './style';

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

  useEffect(() => {
    console.log(user);
  }, [user]);

  return (
    <HeaderWrapper>
      <LeftWrapper>
        <CompanyWrapper>
          <CompanyImage src={`${API_HOST}/api/company/image`} />
          <CompanyName>{data?.company.name}</CompanyName>
        </CompanyWrapper>
      </LeftWrapper>
      <RightWrapper
        onContextMenu={(e) => {
          e.preventDefault();
          console.log(e.clientX, e.clientY);
        }}>
        {/* <img
          alt="User's Profile"
          src={`${API_HOST}/api/profile/image?user-id=${user.id}`}></img> */}
        <ProfileWrapper>
          <ImageWrapper>
            <img
              alt="user's profile"
              src={`https://gravatar.com/avatar/${MD5(
                user.email.trim().toLowerCase(),
              )}?s=36&d=retro`}
            />
          </ImageWrapper>
          <StatusMargin />
          <Status />
        </ProfileWrapper>
      </RightWrapper>
    </HeaderWrapper>
  );
}
