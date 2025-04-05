import { Skeleton } from "@/components/ui/skeleton";

export function EditSizeFormSkeleton() {
  return (
    <div className="space-y-6">
      <div className="grid grid-cols-3 gap-8">
        <div className="space-y-1">
          <Skeleton className="h-4 w-32" />
          <Skeleton className="h-10 w-full" />
        </div>
        <div className="space-y-1">
          <Skeleton className="h-4 w-32" />
          <Skeleton className="h-10 w-full" />
        </div>
      </div>
      <Skeleton className="h-10 w-40 mt-6" />
    </div>
  );
}
