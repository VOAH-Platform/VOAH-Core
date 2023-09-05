import { useEffect } from 'react';
import { useAlert } from 'react-alert';
import { useNavigate, useSearchParams } from 'react-router-dom';

export function VerifyPage() {
  const alert = useAlert();
  const navigate = useNavigate();
  const [params] = useSearchParams();

  const type = params.get('type');
  const code = params.get('code');
  const email = params.get('email');

  useEffect(() => {
    if (!code) {
      alert.error('잘못된 접근입니다.');
      return navigate('/');
    }
    switch (type) {
      case 'email':
        break;
      default:
        alert.error('잘못된 접근입니다.');
        return navigate('/');
    }
  }, [params]);

  return <div>register with {email}</div>;
}
