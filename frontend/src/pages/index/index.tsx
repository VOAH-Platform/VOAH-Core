import { AnimatePresence } from 'framer-motion';
import { useAtom } from 'jotai';
import { Sun, Moon } from 'lucide-react';
import { useEffect, useState } from 'react';
import { useAlert } from 'react-alert';
import { useNavigate } from 'react-router-dom';

import LogoBlack from '@/assets/logo-black.svg?react';
import LogoLight from '@/assets/logo-light.svg?react';
import { themeAtom, userAtom } from '@/atom';
import { FormButton } from '@/components/FormButton';
import { FormInput } from '@/components/FormInput';
import { THEME_TOKEN } from '@/constant';

import { AuthBackground } from './Background';
import { useIndexLogic } from './logic';
import {
  Container,
  ContainerTitle,
  IndexWrapper,
  ContainerHead,
  ContainerBody,
  ButtonWrapper,
  SmallAction,
  ThemeButton,
  AnimWrapper,
  FormError,
} from './style';

enum STEP {
  LOGIN,
  REGISTER,
  PW_RESET,
}

export function IndexPage() {
  const alert = useAlert();
  const navigate = useNavigate();

  const [theme, setTheme] = useAtom(themeAtom);
  const [user, setUser] = useAtom(userAtom);

  const [title, setTitle] = useState('로그인');
  const [step, setStep] = useState(STEP.LOGIN);

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const [formDisabled, setFormDisabled] = useState(false);

  const [error, setError] = useState('');

  const { handleLoginSubmit, handleRegisterSubmit, handlePwResetSubmit } =
    useIndexLogic();

  useEffect(() => {
    if (user.isLogin) {
      return navigate('/app');
    }
  }, [user]);

  const changeStep = (step: STEP) => {
    setStep(step);
    setError('');
    switch (step) {
      case STEP.LOGIN:
        setTitle('로그인');
        break;
      case STEP.REGISTER:
        setTitle('회원가입');
        break;
      case STEP.PW_RESET:
        setTitle('비밀번호 재설정');
        break;
    }
  };

  return (
    <IndexWrapper>
      <Container layout>
        <ContainerHead layout>
          {theme.isDark ? <LogoBlack /> : <LogoLight />}
          <ContainerTitle>{title}</ContainerTitle>
        </ContainerHead>
        <ContainerBody layout>
          <AnimatePresence>
            {step == STEP.LOGIN && (
              <AnimWrapper
                key="login"
                onSubmit={(e) => {
                  e.preventDefault();
                  setFormDisabled(true);
                  void handleLoginSubmit(email, password).then((val) => {
                    console.log(val);
                    setFormDisabled(false);
                    if (!val.success) return setError(val.error);
                    setUser({
                      email: val.value.email,
                      isLogin: true,
                      id: val.value.userId,
                      accessToken: val.value.accessToken,
                      refreshToken: val.value.refreshToken,
                    });
                    return navigate('/app');
                  });
                }}
                layout>
                <FormInput
                  id="login-id"
                  label="이메일"
                  name="login-id"
                  type="text"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  placeholder="이메일을 입력해주세요"
                />
                <FormInput
                  id="login-pw"
                  label="비밀번호"
                  name="login-pw"
                  type="password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  placeholder="비밀번호를 입력해주세요"
                />
                {error !== '' && <FormError layout>{error}</FormError>}
                <ButtonWrapper layout>
                  <FormButton filled type="submit" disabled={formDisabled}>
                    로그인
                  </FormButton>
                  <FormButton
                    type="button"
                    onClick={() => changeStep(STEP.REGISTER)}
                    disabled={formDisabled}>
                    회원가입
                  </FormButton>
                </ButtonWrapper>
                <SmallAction onClick={() => changeStep(STEP.PW_RESET)}>
                  비밀번호 재설정하기
                </SmallAction>
              </AnimWrapper>
            )}
            {step == STEP.REGISTER && (
              <AnimWrapper
                key="register"
                onSubmit={(e) => {
                  e.preventDefault();
                  setFormDisabled(true);
                  void handleRegisterSubmit(email).then((val) => {
                    if (!val.success) setError(val.error);
                    else {
                      alert.success(
                        '회원가입 메일을 보냈습니다. 30분 내로 이메일 인증을 진행해주세요.',
                      );
                      changeStep(STEP.LOGIN);
                    }
                    setFormDisabled(false);
                    return;
                  });
                }}
                layout>
                <FormInput
                  id="register-id"
                  label="이메일"
                  type="email"
                  value={email}
                  onChange={(e) => {
                    setEmail(e.target.value);
                  }}
                  placeholder="이메일을 입력해주세요"
                />
                {error !== '' && <FormError layout>{error}</FormError>}
                <ButtonWrapper layout>
                  <FormButton filled type="submit" disabled={formDisabled}>
                    인증 메일 보내기
                  </FormButton>
                  <FormButton
                    type="button"
                    onClick={() => changeStep(STEP.LOGIN)}
                    disabled={formDisabled}>
                    돌아가기
                  </FormButton>
                </ButtonWrapper>
              </AnimWrapper>
            )}
            {step == STEP.PW_RESET && (
              <AnimWrapper
                key="pw-reset"
                onSubmit={(e) => {
                  e.preventDefault();
                  setFormDisabled(true);
                  void handlePwResetSubmit(email).then((val) => {
                    if (!val.success) setError(val.error);
                    else {
                      alert.success(
                        '비밀번호 초기화 메일을 보냈습니다. 메일함을 확인해주세요',
                      );
                      changeStep(STEP.LOGIN);
                    }
                    setFormDisabled(false);
                    return;
                  });
                }}
                layout>
                <FormInput
                  id="register-id"
                  label="이메일"
                  type="email"
                  value={email}
                  onChange={(e) => {
                    setEmail(e.target.value);
                  }}
                  placeholder="이메일을 입력해주세요"
                />
                {error !== '' && <FormError layout>{error}</FormError>}
                <ButtonWrapper layout>
                  <FormButton filled type="submit" disabled={formDisabled}>
                    인증 메일 보내기
                  </FormButton>
                  <FormButton
                    type="button"
                    onClick={() => changeStep(STEP.LOGIN)}
                    disabled={formDisabled}>
                    돌아가기
                  </FormButton>
                </ButtonWrapper>
              </AnimWrapper>
            )}
          </AnimatePresence>
        </ContainerBody>
      </Container>
      <ThemeButton
        onClick={() => {
          setTheme({
            token: theme.isDark ? THEME_TOKEN.LIGHT : THEME_TOKEN.DARK,
            isDark: !theme.isDark,
          });
        }}>
        {theme.isDark ? <Moon /> : <Sun />}
      </ThemeButton>
      <AuthBackground />
    </IndexWrapper>
  );
}
