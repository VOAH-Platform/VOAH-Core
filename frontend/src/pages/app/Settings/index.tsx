import { useAtom } from 'jotai';
import { SettingsIcon } from 'lucide-react';
import { useEffect } from 'react';
import { Outlet, useLocation, useNavigate } from 'react-router-dom';

import { headerAtom } from '@/atom';

export function VoahSettingsPage() {
  const [, setHeaderData] = useAtom(headerAtom);

  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    setHeaderData({
      isHidden: false,
      icon: <SettingsIcon size={20} />,
      name: '설정',
    });
    navigate('/app/settings/profile');
  }, [location]);

  return (
    <>
      <Outlet />
    </>
  );
}
