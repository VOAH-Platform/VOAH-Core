import { AtomIcon, Building2Icon } from 'lucide-react';
import { useState } from 'react';

import { VoahSideCategoryTypeButton, VoahSideCategoryWrapper } from './style';

export function VoahSideCategory() {
  const [sideType, setSideType] = useState<'company' | 'project'>('company');

  return (
    <VoahSideCategoryWrapper>
      <VoahSideCategoryTypeButton>
        {sideType === 'company' ? <Building2Icon /> : <AtomIcon />}
      </VoahSideCategoryTypeButton>
    </VoahSideCategoryWrapper>
  );
}
