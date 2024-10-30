import { json, MetaFunction } from "@remix-run/node";
import {
  dehydrate,
  HydrationBoundary,
  QueryClient,
  useQuery,
} from "@tanstack/react-query";
import { useLoaderData } from "@remix-run/react";

export const meta: MetaFunction = () => {
  return [
    { title: "Test" },
    { name: "description", content: "Welcome to Remix!" },
  ];
};

const getData = () => ({
  userId: 1,
  id: 1,
  title: "delectus aut autem",
  completed: false,
});

export async function loader() {
  const queryClient = new QueryClient();

  await queryClient.prefetchQuery({
    queryKey: ["data"],
    queryFn: getData,
  });

  return json({ dehydratedState: dehydrate(queryClient) });
}

const Test = () => {
  const { data } = useQuery({ queryKey: ["data"], queryFn: getData });

  return (
    <div className="flex flex-col">
      <span>{data?.userId}</span>
      <span>{data?.title}</span>
      <span>{data?.id}</span>
      <span>{data?.completed}</span>
    </div>
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
