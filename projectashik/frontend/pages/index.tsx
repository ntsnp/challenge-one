import type { NextPage } from "next";
import { useEffect, useState } from "react";

const Index: NextPage = () => {
  const [data, setData] = useState<any>(null);
  const [page, setPage] = useState(1);
  const fetchContents = async (page: number) => {
    const res = await fetch("/api/get-blogs?page=" + page);
    const data = await res.json();
    setData(data.posts);
  };
  useEffect(() => {
    fetchContents(page);
  }, [page]);
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
      }}
    >
      {data &&
        data.map((post: any) => {
          return (
            <a
              style={{
                color: "blue",
                textDecoration: "underline",
              }}
              href={"https://blog.sentry.io" + post.link}
              key={post._id}
            >
              <h1
                style={{
                  textAlign: "center",
                }}
              >
                {post.title}
              </h1>
            </a>
          );
        })}

      <div>
        <button
          style={{
            marginTop: "20px",
            padding: "10px 20px",
            fontSize: "22px",
          }}
          disabled={page === 1}
          onClick={() => setPage(page - 1)}
        >
          Previous
        </button>
        <button
          style={{
            marginTop: "20px",
            marginLeft: "10px",
            padding: "10px 20px",
            fontSize: "22px",
          }}
          onClick={() => setPage(page + 1)}
        >
          Next
        </button>
      </div>
    </div>
  );
};

export default Index;
