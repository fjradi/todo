import { json, LoaderFunctionArgs, MetaFunction } from "@remix-run/node";
import {
  dehydrate,
  HydrationBoundary,
  QueryClient,
  useQuery,
} from "@tanstack/react-query";
import { useLoaderData, useSearchParams } from "@remix-run/react";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  CommandSeparator,
} from "~/components/ui/command";
import { clsx } from "clsx";

export const meta: MetaFunction = () => {
  return [
    { title: "Test" },
    { name: "description", content: "Welcome to Remix!" },
  ];
};

type Todo = {
  id: string;
  name: string;
  is_completed: boolean;
};

const getTodos = async (search: string) => {
  const resp = await fetch(
    `${process.env["BACKEND_URL"]}/todo?search=${search}`
  );
  const data = await resp.json();
  return data as Todo[];
};

export async function loader({ request }: LoaderFunctionArgs) {
  const queryClient = new QueryClient();
  const search = new URL(request.url).searchParams.get("search") ?? "";

  await queryClient.prefetchQuery({
    queryKey: ["todo", { search }],
    queryFn: () => getTodos(search),
  });

  return json({ dehydratedState: dehydrate(queryClient) });
}

const Test = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const search = searchParams.get("search") ?? "";

  const { data } = useQuery({
    queryKey: ["todo", { search }],
    queryFn: () => getTodos(search),
  });

  const handleSearch = (search: string) =>
    setSearchParams((prev) => {
      prev.set("search", search);
      return prev;
    });

  return (
    <Command>
      <CommandInput
        placeholder="Search todos..."
        value={search}
        onValueChange={handleSearch}
      />
      <CommandList>
        <CommandEmpty>No results found.</CommandEmpty>
        <CommandGroup heading="Add todo">
          <CommandItem>{search}</CommandItem>
        </CommandGroup>
        <CommandSeparator />
        <CommandGroup heading="Found todos">
          {data?.map((todo) => (
            <CommandItem
              key={todo.id}
              className={clsx(todo.is_completed && "line-through")}
            >
              {todo.name}
            </CommandItem>
          ))}
        </CommandGroup>
      </CommandList>
    </Command>
  );
};

export default function TestRoute() {
  const { dehydratedState } = useLoaderData<typeof loader>();
  return (
    <HydrationBoundary state={dehydratedState}>
      <Test />
    </HydrationBoundary>
  );
}
