import { Copy, Server } from "lucide-react";
import { Alert, AlertDescription, AlertTitle } from "./ui/alert";
import { Badge } from "@/components/ui/badge";
import toast from "react-hot-toast";
import { Button } from "./ui/button";

interface ApiAlertProps {
  title: string;
  description: string;
  variant?: "public" | "admin";
}

const textMap: Record<NonNullable<ApiAlertProps["variant"]>, string> = {
  public: "Público",
  admin: "Admin",
};

const variantMap: Record<
  NonNullable<ApiAlertProps["variant"]>,
  "secondary" | "destructive"
> = {
  public: "secondary",
  admin: "destructive",
};

export function ApiAlert({
  title,
  description,
  variant = "public",
}: ApiAlertProps) {
  function handleCopy() {
    navigator.clipboard.writeText(description);
    toast.success("rota da api copiado para a área de transferência.");
  }

  return (
    <Alert>
      <Server className="h-4 mr-4" />
      <AlertTitle className="flex items-center gap-x-2">
        {title}
        <Badge variant={variantMap[variant]}>{textMap[variant]}</Badge>
      </AlertTitle>

      <AlertDescription className="mt-4 flex items-center justify-between">
        <code className="relative rounded bg-muted px-[0.3rem] py-[0.2rem] font-mono text-sm font-semibold">
          {description}
        </code>
        <Button variant="outline" size="icon" onClick={handleCopy}>
          <Copy className="h-4 w-4" />
        </Button>
      </AlertDescription>
    </Alert>
  );
}
