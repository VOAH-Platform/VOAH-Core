import { useAtom } from 'jotai';
import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

import { userAtom } from '@/atom';

export function LogoutPage() {
  const [, setUser] = useAtom(userAtom);

  const navigate = useNavigate();

  useEffect(() => {
    setUser({
      email: '',
      isLogin: false,
      id: '',
      accessToken: '',
      refreshToken: '',
    });
    navigate('/');
  });

  return <span>로그아웃하는 중입니다...</span>;
}
