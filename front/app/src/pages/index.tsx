import { Button } from "@chakra-ui/react"
import useSWR from "swr"
import { Fetcher } from "../component/utils"

function Home() {
  const {data} = useSWR("http://localhost:8000", Fetcher)
  if (!data) {
    return <p>no data found...</p>
  }
  return <Button>{data.name}</Button>
}

export default Home