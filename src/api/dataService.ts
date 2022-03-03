import instance from "@/api/instace";

class DataService {
  getAll(): unknown {
    return instance.get("/posts");
  };

  newsLetter(payload: unknown): unknown {
    return instance.post("/posts/email", payload);
  }
}

export default DataService;
