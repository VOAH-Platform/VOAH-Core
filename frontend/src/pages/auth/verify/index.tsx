import { useEffect } from 'react';
import { useAlert } from 'react-alert';
import { useNavigate, useParams } from 'react-router-dom';

export function VerifyPage() {
  const alert = useAlert();
  const navigate = useNavigate();
  const { type, email, code } = useParams();

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
  }, [type, code]);

  return <div>register</div>;
}
