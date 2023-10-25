import { useAtom } from 'jotai';
import { useEffect } from 'react';
import { Outlet, useLocation, useNavigate } from 'react-router-dom';

import { localDataAtom, userAtom } from '@/atom';

import { AppHeader } from './AppHeader';
import { AppWrapper } from './style';
import { VoahFrame } from './VoahFrame';
import { VoahSidebar } from './VoahSidebar';

export function AppLayout() {
  const [user] = useAtom(userAtom);
  const [localData, setLocalData] = useAtom(localDataAtom);

  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    if (!user.isLogin) {
      return navigate('/');
    }
  }, [user]);

  useEffect(() => {
    setLocalData((prev) => {
      return {
        ...prev,
        lastPath: location.pathname,
      };
    });
  }, [location]);

  useEffect(() => {
    if (location.pathname === '/app' && localData.lastPath !== '/app') {
      navigate(localData.lastPath);
    }
  }, []);

  return (
    <>
      <AppHeader />
      <AppWrapper>
        <VoahSidebar />
        {location.pathname.split('/')[2] !== 'm' ? <Outlet /> : <VoahFrame />}
      </AppWrapper>
    </>
  );
}
