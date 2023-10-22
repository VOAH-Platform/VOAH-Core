import { icons } from 'lucide-react';

export function LucideCustom({
  icon,
  size,
  color,
  ...props
}: {
  icon: string;
  size: number;
  color: string;
}) {
  const Icon = icons[icon as keyof typeof icons];

  if (!Icon) {
    throw new Error(`Invalid icon name: ${icon}`);
  }

  return <Icon size={size} color={color} {...props} />;
}
