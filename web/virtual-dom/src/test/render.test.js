import render from "../vdom/render";

describe("render", () => {
  test("renders a text node", () => {
    const result = render("Hello, World!");
    expect(result).toBeInstanceOf(Text);
    expect(result.textContent).toBe("Hello, World!");
  });

  test("renders an element with attributes and children", () => {
    const vNode = {
      tagName: "div",
      attrs: { id: "test", class: "container" },
      children: ["Child 1", "Child 2"],
    };
    const result = render(vNode);

    expect(result.tagName).toBe("DIV");
    expect(result.getAttribute("id")).toBe("test");
    expect(result.getAttribute("class")).toBe("container");
    expect(result.childNodes.length).toBe(2);
    expect(result.childNodes[0].textContent).toBe("Child 1");
    expect(result.childNodes[1].textContent).toBe("Child 2");
  });
});
