import { useOrigin } from "@/hooks/use-origin";
import { useParams } from "react-router-dom";
import { ApiAlert } from "./api-alert";

interface ApiListProps {
  entityName: string;
  entityIdName: string;
}

export function ApiList({ entityName, entityIdName }: ApiListProps) {
  const { storeId } = useParams();
  const origin = useOrigin();

  const baseUrl = `${origin}/api/v1/stores/${storeId}/${entityName}`;

  return (
    <>
      <ApiAlert title="GET" description={baseUrl} />
      <ApiAlert title="GET" description={`${baseUrl}/{${entityIdName}}`} />
      <ApiAlert title="POST" variant="admin" description={baseUrl} />
      <ApiAlert
        title="PUT"
        variant="admin"
        description={`${baseUrl}/{${entityIdName}}`}
      />
      <ApiAlert
        title="DELETE"
        variant="admin"
        description={`${baseUrl}/{${entityIdName}}`}
      />
    </>
  );
}
