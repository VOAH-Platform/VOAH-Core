import { useQuery } from '@tanstack/react-query';
import { useAtom } from 'jotai';
import { FileSymlink, Moon, Sun } from 'lucide-react';
import { useEffect, useState } from 'react';
import { useAlert } from 'react-alert';
import { Link, useNavigate, useSearchParams } from 'react-router-dom';

import { apiClient } from '@/apiClient';
import LogoWithText from '@/assets/logo-with-text.svg?react';
import { themeAtom } from '@/atom';
import { FormButton } from '@/components/FormButton';
import { FormCheckbox } from '@/components/FormCheckbox';
import { FormInput } from '@/components/FormInput';
import { FormSelect } from '@/components/FormSelect';
import { THEME_TOKEN } from '@/constant';
import { css } from '@/stitches.config';

import { useVerifyLogic } from './logic';
import {
  RegisterBody,
  RegisterForm,
  RegisterHeader,
  RegisterTitle,
  RegisterWrapper,
  ThemeButton,
  TosLinkWrapper,
  TosTitle,
  TosWrapper,
} from './style';

interface VerifyData {
  message: string;
  'public-teams': {
    'team-id': string;
    public: boolean;
    displayname: string;
    description: string;
    'created-at': string;
  }[];
  'invited-teams': {
    'team-id': string;
    public: boolean;
    displayname: string;
    description: string;
    'created-at': string;
  }[];
}

export function VerifyPage() {
  const alert = useAlert();
  const navigate = useNavigate();

  const [theme, setTheme] = useAtom(themeAtom);

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [pwcheck, setPwcheck] = useState('');
  const [username, setUsername] = useState('');
  const [nickname, setNickname] = useState('');
  const [position, setPosition] = useState('');
  const [team, setTeam] = useState('');
  const [tos, setTos] = useState(false);

  const [params] = useSearchParams();

  const { emailVerify } = useVerifyLogic();

  const paramType = params.get('type');
  const paramCode = params.get('code');
  const paramEmail = params.get('email');

  useEffect(() => {
    if (!paramCode) {
      alert.error('잘못된 접근입니다.');
      return navigate('/');
    }
    switch (paramType) {
      case 'email':
        if (!paramEmail) {
          alert.error('잘못된 접근입니다.');
          return navigate('/');
        }
        setEmail(paramEmail);
        break;
      default:
        alert.error('잘못된 접근입니다.');
        return navigate('/');
    }
  }, [params]);

  const { isLoading, isError, data } = useQuery<unknown, unknown, VerifyData>({
    queryKey: ['verify'],
    enabled: paramCode !== '',
    queryFn: async () => {
      try {
        const data = await apiClient.post<VerifyData>(
          '/api/auth/register/check',
          {
            email: paramEmail,
            code: paramCode,
          },
        );
        return data.data;
      } catch (e) {
        console.log(e);
        throw e;
      }
    },
  });

  const handleFormSubmit = async () => {
    switch (paramType) {
      case 'email': {
        const res = await emailVerify({
          code: paramCode!,
          email,
          password,
          pwcheck,
          username,
          nickname,
          position,
          team,
          tos,
        });
        if (res.success) {
          alert.success('회원가입이 완료되었습니다. 로그인을 진행해주세요.');
          return navigate('/');
        } else {
          alert.error(res.error);
          // TODO: 나중에 폼에 직접 표시되는 에러로 수정 필요
        }
        break;
      }
    }
  };

  return (
    <RegisterWrapper>
      <RegisterHeader>
        <LogoWithText />
        <ThemeButton
          onClick={() => {
            setTheme({
              token: theme.isDark ? THEME_TOKEN.LIGHT : THEME_TOKEN.DARK,
              isDark: !theme.isDark,
            });
          }}>
          {theme.isDark ? <Moon /> : <Sun />}
        </ThemeButton>
      </RegisterHeader>
      {isLoading && <RegisterBody>정보를 불러오고 있습니다...</RegisterBody>}
      {isError && (
        <RegisterBody>
          올바르지 않은 요청입니다.
          <br />
          <Link to="/">처음으로 돌아가기</Link>
        </RegisterBody>
      )}
      {data && !isError && (
        <RegisterBody>
          <RegisterTitle>회원가입</RegisterTitle>
          <RegisterForm
            onSubmit={(e) => {
              e.preventDefault();
              void handleFormSubmit();
            }}>
            <FormInput
              id="register-id"
              label="이메일"
              type="email"
              placeholder="이메일을 입력해주세요"
              value={email}
              onChange={(e) => {
                setEmail(e.target.value);
              }}
              disabled={paramEmail !== ''}
            />
            <FormInput
              id="register-password"
              label="비밀번호"
              type="password"
              placeholder="비밀번호를 입력해주세요"
              value={password}
              onChange={(e) => {
                setPassword(e.target.value);
              }}
            />
            <FormInput
              id="register-pwcheck"
              label="비밀번호 확인"
              type="password"
              placeholder="비밀번호를 다시 입력해주세요"
              value={pwcheck}
              onChange={(e) => {
                setPwcheck(e.target.value);
              }}
            />
            <FormInput
              id="register-username"
              label="사용자 이름"
              type="text"
              placeholder="사용자 이름을 입력해주세요"
              value={username}
              onChange={(e) => {
                setUsername(e.target.value);
              }}
            />
            <FormInput
              id="register-nickname"
              label="별명"
              type="text"
              placeholder="별명을 입력해주세요"
              value={nickname}
              onChange={(e) => {
                setNickname(e.target.value);
              }}
            />
            <FormInput
              id="register-position"
              label="직책"
              type="text"
              placeholder="직책을 입력해주세요"
              value={position}
              onChange={(e) => {
                setPosition(e.target.value);
              }}
            />
            <FormSelect
              id="register-team"
              label="팀 선택"
              items={data['invited-teams']
                .concat(data['public-teams'])
                .map((team) => ({
                  value: team['team-id'],
                  label: team.displayname,
                  private: !team.public,
                }))}
              placeholder="가입할 팀을 선택해주세요"
              onSelect={(val) => {
                setTeam(val.value);
              }}
            />
            <TosWrapper>
              <TosTitle>이용약관</TosTitle>
              <Link to="https://naver.com" target="_blank">
                <TosLinkWrapper>
                  <FileSymlink />
                  이용약관 보기
                </TosLinkWrapper>
              </Link>
              <FormCheckbox
                id="register-tos"
                text="이용약관에 동의합니다"
                onChange={(e) => {
                  setTos(e.target.checked);
                }}
              />
            </TosWrapper>
            <FormButton
              className={css({
                marginBottom: '2rem',
              })()}
              type="submit"
              filled>
              회원가입
            </FormButton>
          </RegisterForm>
        </RegisterBody>
      )}
    </RegisterWrapper>
  );
}
