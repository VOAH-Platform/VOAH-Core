import { VoahSideCategory } from './Category';
import { VoahSideMenu } from './Menu/component';
import { VoahSideWrapper } from './style';

export function VoahSidebar() {
  return (
    <VoahSideWrapper>
      <VoahSideCategory />
      <VoahSideMenu />
    </VoahSideWrapper>
  );
}
