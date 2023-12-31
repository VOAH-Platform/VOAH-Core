import { styled } from '@/stitches.config';

export const HeaderWrapper = styled('header', {
  padding: '0 1.5rem',
  width: '100%',
  height: '60px',
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
  boxShadow: '$grade1',
});

export const LeftWrapper = styled('div', {
  display: 'flex',
  alignItems: 'center',
});

export const CompanyWrapper = styled('div', {
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
  gap: '1rem',
});

export const CompanyImage = styled('img', {
  width: '36px',
  height: '36px',
  borderRadius: '0.75rem',
  background: '$gray100',
});

export const CompanyName = styled('span', {
  overflow: 'hidden',
  color: '$gray900',
  textOverflow: 'ellipsis',
  fontSize: '20px',
  fontStyle: 'normal',
  fontWeight: '700',
  lineHeight: '140%',
  letterSpacing: '-0.2px',
});

export const TitleDivider = styled('div', {
  width: '1px',
  height: '24px',
  background: '$gray300',
  margin: '0 0.5rem',
});

export const TitleWrapper = styled('div', {
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
  gap: '0.5rem',
});

export const Title = styled('span', {
  overflow: 'hidden',
  color: '$gray900',
  textOverflow: 'ellipsis',
  fontSize: '1.25rem',
  fontStyle: 'normal',
  fontWeight: 600,
  letterSpacing: '-0.2px',
});

export const RightWrapper = styled('div', {
  display: 'flex',
  alignItems: 'center',
});

export const ProfileWrapper = styled('div', {
  width: '36px',
  aspectRatio: '1/1',
  overflow: 'hidden',
  background: '$gray0',
  position: 'relative',
  cursor: 'pointer',
});

export const ImageWrapper = styled('div', {
  width: '36px',
  aspectRatio: '1/1',
  overflow: 'hidden',
  borderRadius: '50%',
  background: '$gray100',
});

export const StatusMargin = styled('div', {
  width: '19.5px',
  height: '19.5px',
  position: 'absolute',
  borderRadius: '50%',
  bottom: '-3px',
  right: '-3px',
  background: '$gray0',
});

export const Status = styled('div', {
  width: '13.5px',
  height: '13.5px',
  position: 'absolute',
  borderRadius: '50%',
  bottom: 0,
  right: 0,
  background: '$accept500',
});
