import { useAtom } from 'jotai';
import { HandIcon, KeyRoundIcon } from 'lucide-react';
import { useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

import { API_HOST } from '@/apiClient';
import { userAtom } from '@/atom';

import { useSideMenu } from '../VoahSidebar/Menu';

export function useSettingsInit() {
  const [user] = useAtom(userAtom);

  const sideMenu = useSideMenu();

  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    sideMenu.setSideMenu(
      {
        title: 'VOAH 설정',
        desc: 'VOAH 설정 페이지',
        hideDesc: true,
      },
      [],
      [
        {
          icon: (
            <img
              width="20"
              style={{
                width: '20px',
                height: '20px',
                borderRadius: '50%',
                overflow: 'hidden',
              }}
              alt="User's Profile"
              src={`${API_HOST}/api/profile/image?user-id=${(
                JSON.parse(localStorage.getItem('voah__user')!) as {
                  id: string;
                }
              )?.id}`}
            />
          ),
          name: '프로필',
          onClick: () => {
            navigate('/app/settings/profile');
          },
          subMenu: [],
        },
        {
          icon: <KeyRoundIcon size={20} />,
          name: '개인정보 및 보안',
          onClick: () => {
            alert('준비중입니다.');
          },
          subMenu: [],
        },
        {
          icon: <HandIcon size={20} />,
          name: '테마 및 접근성',
          onClick: () => {
            navigate('/app/settings/accessibility');
          },
          subMenu: [],
        },
      ],
    );
  }, [user.id, location]);
}
