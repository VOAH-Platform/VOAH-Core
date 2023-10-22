import { useQuery } from '@tanstack/react-query';
import { useAtom } from 'jotai';
import { AtomIcon, Building2Icon, MessageSquareIcon } from 'lucide-react';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { userAtom } from '@/atom';
import { getProfileById } from '@/lib/query/profile';

import { VoahSideCategoryTypeButton, VoahSideCategoryWrapper } from './style';

export function VoahSideCategory() {
  const [user] = useAtom(userAtom);

  const [sideType, setSideType] = useState<'company' | 'project'>('company');

  const navigate = useNavigate();

  const { data } = useQuery({
    queryKey: ['myData'],
    queryFn: async () => {
      return await getProfileById(user.id, user.accessToken);
    },
    enabled: user.isLogin,
  });

  return (
    <VoahSideCategoryWrapper>
      <VoahSideCategoryTypeButton
        onClick={() => {
          if (
            sideType === 'company' &&
            (!data ||
              !data.data.user.projects ||
              data?.data.user.projects.length === 0)
          ) {
            alert('프로젝트에 소속되어 있지 않습니다.');
          } else {
            setSideType((prev) => (prev === 'company' ? 'project' : 'company'));
          }
        }}>
        {sideType === 'company' ? (
          <Building2Icon size={28} />
        ) : (
          <AtomIcon size={28} />
        )}
      </VoahSideCategoryTypeButton>
      {sideType === 'company' && (
        <>
          <VoahSideCategoryTypeButton
            onClick={() => {
              navigate('/app/m');
            }}>
            <MessageSquareIcon size={24} />
          </VoahSideCategoryTypeButton>
        </>
      )}
      {sideType === 'project' && <></>}
    </VoahSideCategoryWrapper>
  );
}
