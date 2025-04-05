import { Button, buttonVariants } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { Loader2 } from "lucide-react";
import { VariantProps } from "class-variance-authority";

interface LoadingButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof buttonVariants> {
  isLoading?: boolean;
}

export function LoadingButton({
  isLoading,
  children,
  className,
  variant,
  size,
  ...props
}: LoadingButtonProps) {
  return (
    <Button
      className={cn(
        "relative flex items-center justify-center overflow-hidden",
        buttonVariants({ variant, size }),
        className
      )}
      disabled={isLoading}
      {...props}
    >
      <div className="relative flex items-center justify-center w-full">
        <span
          className={cn(
            "transition-all duration-300",
            isLoading
              ? "-translate-y-full opacity-0"
              : "translate-y-0 opacity-100"
          )}
        >
          {children}
        </span>

        <span
          className={cn(
            "absolute flex items-center justify-center transition-all duration-300",
            isLoading
              ? "translate-y-0 opacity-100"
              : "translate-y-full opacity-0"
          )}
        >
          <Loader2 className="animate-spin h-5 w-5" />
        </span>
      </div>
    </Button>
  );
}
