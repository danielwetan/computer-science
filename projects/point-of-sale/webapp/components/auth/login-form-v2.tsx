"use client";

import Link from "next/link";
import { LoginSchema } from "@/lib/validators/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { FormProvider, useForm } from "react-hook-form";
import { loginQuery } from "@/queries/auth";
import { z } from "zod";
import { jwtDecode } from "jwt-decode";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";

const LoginForm = () => {
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
        const decodedToken: any = jwtDecode(response?.data?.jwt_token);
        console.log(decodedToken);
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
