import { Skeleton } from "@/components/ui/skeleton";

export function EditCategoryFormSkeleton() {
  return (
    <div className="space-y-8 w-full">
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

      <Skeleton className="h-10 w-40" />
    </div>
  );
}
