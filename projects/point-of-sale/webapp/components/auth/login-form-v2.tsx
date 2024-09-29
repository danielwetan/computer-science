"use client";

import { LoginSchema } from "@/lib/validators/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { loginQuery } from "@/queries/auth";
import { profileQuery } from "@/queries/profile";
import { z } from "zod";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useAppDispatch, useAppSelector } from "@/stores";
import { setData } from "@/stores/slices/auth";
import { setDialog } from "@/stores/slices/dialog";
import { createSession } from "@/lib/session";

const LoginForm = () => {
  const {
    data: { login },
  } = useAppSelector((state) => state.dialogReducer);
  const dispatch = useAppDispatch();

  const router = useRouter();

  const BaseLoginSchema = LoginSchema();
  const form = useForm<z.infer<typeof BaseLoginSchema>>({
    resolver: zodResolver(BaseLoginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof BaseLoginSchema>) => {
    try {
      const response = await loginQuery({
        email: values?.email,
        password: values?.password,
      });
      if (response?.status === 200) {
        console.log("trying to store session");
        await createSession(response?.data?.jwt_token);
      }
      const responseProfile = await profileQuery();
      if (responseProfile.status === 200) {
        dispatch(
          setData({
            data: responseProfile?.data,
          })
        );
        dispatch(setDialog({ login: false }));
        router.push("/dashboard");
      }
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input placeholder="email" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <Input placeholder="password" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit">Submit</Button>
      </form>
    </Form>
  );
};

export default LoginForm;
