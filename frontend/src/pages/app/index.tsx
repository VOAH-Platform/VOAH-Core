import { useAtom } from 'jotai';
import { useEffect } from 'react';
import { Outlet, useNavigate } from 'react-router-dom';

import { userAtom } from '@/atom';

import { AppHeader } from './AppHeader';
import { VoahFrame } from './VoahFrame';

export function AppLayout() {
  const [user] = useAtom(userAtom);

  const navigate = useNavigate();

  useEffect(() => {
    if (!user.isLogin) {
      return navigate('/');
    }
  }, [user]);

  return (
    <>
      <AppHeader />
      <Outlet />
      <VoahFrame />
    </>
  );
}
