import { useQuery } from '@tanstack/react-query';
import { format } from 'date-fns';
import { useAtom } from 'jotai';
import { User2Icon, UserCog2Icon } from 'lucide-react';

import { API_HOST } from '@/apiClient';
import { userAtom } from '@/atom';
import { getProfileById } from '@/lib/query/profile';

import {
  VoahSettingsProfileEditButton,
  VoahSettingsProfileHead,
  VoahSettingsProfileOtherTitle,
  VoahSettingsProfileOtherTitleWrapper,
  VoahSettingsProfileOtherValue,
  VoahSettingsProfileOtherWrapper,
  VoahSettingsProfileUserInfoDisplayName,
  VoahSettingsProfileUserInfoImage,
  VoahSettingsProfileUserInfoTextWrapper,
  VoahSettingsProfileUserInfoUsername,
  VoahSettingsProfileUserInfoWrapper,
  VoahSettingsProfileWrapper,
} from './style';

export function VoahSettingsProfile() {
  const [user] = useAtom(userAtom);

  const { data } = useQuery({
    queryKey: ['myData'],
    queryFn: async () => {
      return await getProfileById(user.id, user.accessToken);
    },
    enabled: user.isLogin,
  });

  const handleEditButton = () => {
    alert('준비중입니다.');
  };

  return (
    <VoahSettingsProfileWrapper>
      {data && (
        <>
          <VoahSettingsProfileHead>
            <VoahSettingsProfileUserInfoWrapper>
              <VoahSettingsProfileUserInfoImage>
                <img
                  width="128"
                  alt="User's Profile"
                  src={`${API_HOST}/api/profile/image?user-id=${user.id}`}
                />
              </VoahSettingsProfileUserInfoImage>
              <VoahSettingsProfileUserInfoTextWrapper>
                <VoahSettingsProfileUserInfoDisplayName>
                  {data.data.user.displayname}
                </VoahSettingsProfileUserInfoDisplayName>
                <VoahSettingsProfileUserInfoUsername>
                  {data.data.user.username}
                </VoahSettingsProfileUserInfoUsername>
              </VoahSettingsProfileUserInfoTextWrapper>
            </VoahSettingsProfileUserInfoWrapper>
            <VoahSettingsProfileEditButton onClick={handleEditButton}>
              <UserCog2Icon size={20} />
              &nbsp;수정하기
            </VoahSettingsProfileEditButton>
          </VoahSettingsProfileHead>
          <hr />
          <VoahSettingsProfileOtherWrapper>
            <VoahSettingsProfileOtherTitleWrapper>
              <User2Icon size={20} />
              <VoahSettingsProfileOtherTitle>
                포지션
              </VoahSettingsProfileOtherTitle>
            </VoahSettingsProfileOtherTitleWrapper>
            <VoahSettingsProfileOtherValue>
              {data.data.user.position}
            </VoahSettingsProfileOtherValue>
          </VoahSettingsProfileOtherWrapper>
          <hr />
          <VoahSettingsProfileOtherWrapper>
            <VoahSettingsProfileOtherTitleWrapper>
              <User2Icon size={20} />
              <VoahSettingsProfileOtherTitle>
                내 소개
              </VoahSettingsProfileOtherTitle>
            </VoahSettingsProfileOtherTitleWrapper>
            <VoahSettingsProfileOtherValue>
              {data.data.user.description || '소개가 비어있습니다.'}
            </VoahSettingsProfileOtherValue>
          </VoahSettingsProfileOtherWrapper>
          <hr />
          <VoahSettingsProfileOtherWrapper>
            <VoahSettingsProfileOtherTitleWrapper>
              <User2Icon size={20} />
              <VoahSettingsProfileOtherTitle>
                계정 생성일
              </VoahSettingsProfileOtherTitle>
            </VoahSettingsProfileOtherTitleWrapper>
            <VoahSettingsProfileOtherValue>
              {format(
                new Date(data.data.user['created-at']),
                'yyyy년 MM월 dd일',
              )}
            </VoahSettingsProfileOtherValue>
          </VoahSettingsProfileOtherWrapper>
        </>
      )}
    </VoahSettingsProfileWrapper>
  );
}
