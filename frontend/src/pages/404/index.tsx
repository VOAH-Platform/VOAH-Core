import NotFoundLogo from '@/assets/NotFound.svg';

import { NotFoundButton, NotFoundSubtitle, NotFoundWrapper } from './style';

export function NotFoundPage() {
  return (
    <NotFoundWrapper>
      <NotFoundLogo />
      <NotFoundSubtitle>페이지를 찾을 수 없습니다.</NotFoundSubtitle>
      <NotFoundButton href="/">
        &lt;&nbsp;&nbsp;&nbsp;이전 페이지로 돌아가기
      </NotFoundButton>
    </NotFoundWrapper>
  );
}
