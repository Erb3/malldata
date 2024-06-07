import { createLazyFileRoute } from '@tanstack/react-router';

export const Route = createLazyFileRoute('/')({
  component: Index,
});

function Index() {
  return (
    <div className="flex justify-center">
      <main className="w-full lg:w-3/5 flex flex-col items-center [&>h2]:mt-8 [&>h2]:mb-2 [&>h2]:text-xl [&>h2]:w-full">
        <h1>Shops</h1>
      </main>
    </div>
  );
}
