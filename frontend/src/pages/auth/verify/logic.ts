import { apiClient } from '@/apiClient';
import { Result } from '@/types/Result';

export const useVerifyLogic = () => {
  return {
    emailVerify: async (data: {
      code: string;
      email: string;
      password: string;
      pwcheck: string;
      username: string;
      nickname: string;
      position: string;
      team: string;
      tos: boolean;
    }): Promise<Result<boolean>> => {
      if (!data.tos) {
        return {
          success: false,
          error: '이용약관에 동의해주세요.',
        };
      }
      if (data.password !== data.pwcheck) {
        return {
          success: false,
          error: '비밀번호가 일치하지 않습니다.',
        };
      }
      if (new RegExp(/^[a-z0-9_-]{4,30}$/).test(data.username) === false) {
        return {
          success: false,
          error:
            '유저명은 4~30자의 영문 소문자와 숫자, 언더바(_), 하이폰(-)로만 입력해주세요.',
        };
      }
      if (
        new RegExp(/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/).test(
          data.email,
        ) === false
      ) {
        return {
          success: false,
          error: '이메일 형식이 올바르지 않습니다.',
        };
      }
      if (data.nickname.length > 30) {
        return {
          success: false,
          error: '별명은 30자 이하여야 합니다.',
        };
      }
      if (data.nickname === '') {
        data.nickname = data.username;
      }
      if (data.team === '') {
        return {
          success: false,
          error: '가입할 팀을 선택해주세요.',
        };
      }

      try {
        await apiClient.post('/api/auth/register/check', {
          code: data.code,
          email: data.email,
          password: data.password,
          username: data.username,
          displayname: data.nickname,
          position: data.position,
          'team-id': data.team,
        });
        return {
          success: true,
          value: true,
        };
      } catch (e) {
        return {
          success: false,
          error: '회원가입에 실패했습니다.',
        };
      }
    },
  };
};
