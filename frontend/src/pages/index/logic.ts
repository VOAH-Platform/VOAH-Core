import { AxiosError } from 'axios';

import { apiClient } from '@/apiClient';
import { Result } from '@/types/Result';

// create hook
export const useIndexLogic = () => {
  const checkEmailRegex = (email: string): boolean => {
    const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    return emailRegex.test(email);
  };

  return {
    handleLoginSubmit: async (
      id: string,
      password: string,
    ): Promise<
      Result<{
        email: string;
        userId: string;
        accessToken: string;
        refreshToken: string;
      }>
    > => {
      if (!checkEmailRegex(id)) {
        return {
          success: false,
          error: '이메일 형식이 올바르지 않습니다.',
        };
      }
      try {
        // Web     DeviceType = 1
        // Android DeviceType = 2
        // IOS     DeviceType = 3
        // Windows DeviceType = 4
        // MacOS   DeviceType = 5
        // Linux   DeviceType = 6
        const data = await apiClient.post<{
          'access-token': string;
          exp: number;
          message: string;
          'refresh-token': string;
          'user-id': string;
        }>('/api/auth/login', {
          email: id,
          password: password,
          'device-id': crypto.randomUUID(),
          'device-type': 1, // TODO: device typing with native messaging
          'device-detail': `${
            window.navigator.platform
          }/${window.navigator.userAgent.split(' ').at(-1)!}`, // user-agent
        });
        if (data.status === 200) {
          console.log('success');
          return {
            success: true,
            value: {
              email: id,
              userId: data.data['user-id'],
              accessToken: data.data['access-token'],
              refreshToken: data.data['refresh-token'],
            },
          };
        }
      } catch (e) {
        const error = e as AxiosError;
        console.log(error);
        if (error.response?.status === 401) {
          return {
            success: false,
            error: '이메일 또는 비밀번호가 일치하지 않습니다.',
          };
        }
      }
      return {
        success: false,
        error: '로그인에 실패했습니다.',
      };
    },
    handleRegisterSubmit: async (email: string): Promise<Result<boolean>> => {
      if (!checkEmailRegex(email)) {
        return {
          success: false,
          error: '이메일 형식이 올바르지 않습니다.',
        };
      }
      try {
        const data = await apiClient.post('/api/auth/register', {
          email,
        });
        if (data.status === 200) {
          console.log('success');
          return {
            success: true,
            value: true,
          };
        }
      } catch (e) {
        const error = e as AxiosError;
        console.log(error);
        if (error.response?.status === 409) {
          return {
            success: false,
            error: '이미 가입된 이메일입니다.',
          };
        }
        const data = error.response?.data as { error: string; message: string };
        if (
          error.response?.status === 400 &&
          data.error === 'Email is not from allowed domain'
        ) {
          return {
            success: false,
            error: '가입이 허용된 이메일이 아닙니다.',
          };
        }
      }
      return {
        success: false,
        error: '회원가입에 실패했습니다.',
      };
    },
    handlePwResetSubmit: async (email: string): Promise<Result<boolean>> => {
      if (!checkEmailRegex(email)) {
        return {
          success: false,
          error: '이메일 형식이 올바르지 않습니다.',
        };
      }
      try {
        const data = await apiClient.get(`/api/auth/passreset?email=${email}`);
        if (data.status === 200) {
          console.log('success');
          return {
            success: true,
            value: true,
          };
        }
      } catch (e) {
        console.log(e);
      }
      return {
        success: false,
        error: '비밀번호 재설정에 실패했습니다.',
      };
    },
  };
};
