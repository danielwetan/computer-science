import z from "zod";

export const LoginSchema = () =>
  z.object({
    email: z.string().email("Email tidak valid").min(1, {
      message: "Email wajib diisi",
    }),
    password: z.string().min(8, {
      message: "Password wajib diisi",
    }),
  });
