import { useQuery } from '@tanstack/react-query';
import { useAtom } from 'jotai';
import { LogOutIcon, SettingsIcon } from 'lucide-react';
import { useRef } from 'react';
import { useNavigate } from 'react-router-dom';

import { API_HOST, apiClient } from '@/apiClient';
import { headerAtom, userAtom } from '@/atom';
import { useCustomContext } from '@/lib/context';

import {
  CompanyName,
  HeaderWrapper,
  ProfileWrapper,
  LeftWrapper,
  RightWrapper,
  ImageWrapper,
  CompanyWrapper,
  CompanyImage,
  TitleDivider,
  TitleWrapper,
} from './style';

export function AppHeader() {
  const [user] = useAtom(userAtom);
  const [headerData] = useAtom(headerAtom);

  const { showContext, hideContext } = useCustomContext();

  const navigate = useNavigate();

  const profileRef = useRef<HTMLDivElement>(null);

  const { data } = useQuery({
    queryKey: ['user'],
    enabled: !!user.accessToken,
    queryFn: async () => {
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

  const handleProfileClick = (e: React.MouseEvent<HTMLDivElement>) => {
    e.preventDefault();
    const x = profileRef.current?.getBoundingClientRect().x ?? 0;
    const y = profileRef.current?.getBoundingClientRect().y ?? 0;
    showContext(x + 36, y + 44, [
      {
        nameHidden: false,
        id: 'account',
        name: '계정',
        buttons: [
          {
            icon: <SettingsIcon size={20} />,
            name: '설정',
            onClick: () => {
              hideContext();
              navigate('/app/settings/profile');
            },
          },
          {
            icon: <LogOutIcon size={20} />,
            name: '로그아웃',
            onClick: () => {
              hideContext();
              navigate('/auth/logout');
            },
            isRed: true,
          },
        ],
      },
    ]);
  };

  return (
    <HeaderWrapper>
      <LeftWrapper>
        <CompanyWrapper>
          <CompanyImage src={`${API_HOST}/api/company/image`} />
          <CompanyName>{data?.company.name}</CompanyName>
          {!headerData.isHidden && (
            <>
              <TitleDivider />
              <TitleWrapper>
                {headerData.icon}
                <span>{headerData.name}</span>
              </TitleWrapper>
            </>
          )}
        </CompanyWrapper>
      </LeftWrapper>
      <RightWrapper>
        <ProfileWrapper ref={profileRef} onClick={handleProfileClick}>
          <ImageWrapper>
            <img
              width="36"
              alt="User's Profile"
              src={`${API_HOST}/api/profile/image?user-id=${user.id}`}
            />
          </ImageWrapper>
          {/* <StatusMargin />
          <Status /> */}
        </ProfileWrapper>
      </RightWrapper>
    </HeaderWrapper>
  );
}
