interface RecentArticle {
  id: string | number
  title: string,
  time: string | number,
  sharp: string,
  description: string,
  image?: string
}

export default RecentArticle;