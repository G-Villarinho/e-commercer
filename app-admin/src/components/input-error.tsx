import { Info } from "lucide-react";

interface InputErrorProps {
  error?: string;
}

export function InputError({ error }: InputErrorProps) {
  return (
    <small className="flex gap-2 items-center mt-1 text-red-500 text-sm font-medium">
      <Info size={16} />
      {error}
    </small>
  );
}
