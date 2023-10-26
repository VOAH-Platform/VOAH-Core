import { useAtom } from 'jotai';

import { API_HOST } from '@/apiClient';
import { userAtom } from '@/atom';

export function ProfileImage({ width }: { width?: string }) {
  const [user] = useAtom(userAtom);

  return (
    <div>
      <img
        width={width || '128'}
        alt="User's Profile"
        src={`${API_HOST}/api/profile/image?user-id=${user.id}`}
      />
    </div>
  );
}
