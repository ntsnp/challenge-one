import { NextApiRequest, NextApiResponse } from "next";
import { connectToDatabase } from "../../libs/mongodb";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  const postsPerPage = 20;
  const page = +(req.query.page as string) || 0;
  const posts: any = [];
  const { db } = await connectToDatabase();
  // Send all the todos
  db.collection("posts")
    .find()
    .skip(postsPerPage * page)
    .limit(postsPerPage)
    .forEach((post: any) => {
      delete post.content;
      posts.push(post);
    })
    .then(() => {
      return res.json({
        perPage: postsPerPage,
        posts,
      });
    })
    .catch(() => {
      res.json({ error: "Something went wrong" });
    });
}
