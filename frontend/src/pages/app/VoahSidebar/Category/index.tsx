import { useQuery } from '@tanstack/react-query';
import { useAtom } from 'jotai';
import { AtomIcon, Building2Icon, MessageSquareIcon } from 'lucide-react';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { moduleAtom, userAtom } from '@/atom';
import { getProfileById } from '@/lib/query/profile';

import { VoahSideCategoryTypeButton, VoahSideCategoryWrapper } from './style';

export function VoahSideCategory() {
  const [user] = useAtom(userAtom);
  const [module] = useAtom(moduleAtom);

  const [sideType, setSideType] = useState<'company' | 'project'>('company');

  const navigate = useNavigate();

  const { data } = useQuery({
    queryKey: ['myData'],
    queryFn: async () => {
      const data = await getProfileById(user.id, user.accessToken);
      console.log('myData', data);
      return data;
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
          {module.data.map((val, idx) => {
            if (!val.expose) return null;

            return (
              <VoahSideCategoryTypeButton
                key={idx}
                onClick={() => {
                  navigate(`/app/m/${val.id}`);
                }}>
                {val.name === 'VOAH-Official-Message' && (
                  <MessageSquareIcon size={24} />
                )}
              </VoahSideCategoryTypeButton>
            );
          })}
        </>
      )}
      {sideType === 'project' && <></>}
    </VoahSideCategoryWrapper>
  );
}
