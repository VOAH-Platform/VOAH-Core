import { useAtom } from 'jotai';
import { useEffect } from 'react';
import { Outlet, useLocation, useNavigate } from 'react-router-dom';

import { apiClient } from '@/apiClient';
import { localDataAtom, moduleAtom, ModuleData, userAtom } from '@/atom';

import { AppHeader } from './AppHeader';
import { AppWrapper } from './style';
import { VoahFrame } from './VoahFrame';
import { VoahSidebar } from './VoahSidebar';

export function AppLayout() {
  const [user, setUser] = useAtom(userAtom);
  const [localData, setLocalData] = useAtom(localDataAtom);
  const [, setModule] = useAtom(moduleAtom);

  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    if (!user.isLogin) {
      return navigate('/');
    }

    apiClient
      .get('/api/info/modules', {
        headers: {
          Authorization: `Bearer ${user.accessToken}`,
        },
      })
      .then((res) => {
        const data = res.data as {
          modules: Array<ModuleData>;
          success: boolean;
        };

        setModule((prev) => {
          return {
            ...prev,
            data: data.modules,
          };
        });

        data.modules.forEach((val, idx) => {
          setModule((prev2) => {
            return {
              ...prev2,
              indexMap: prev2.indexMap.set(val.id, idx),
            };
          });
        });

        return;
      })
      .catch(() => {
        navigate('/');
        setUser((prev) => ({
          ...prev,
          isLogin: false,
        }));
      });
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
