# Variable shadowing

- A variable declared within a certain scope (like a function or a block) has the same name as a variable declared in an outer scope -> the inner var shadows the outer var -> the inner var is accessible in that scope -> the outer var is temporarily hidden

## Key points

1. **Scope**: Shadowing is about variable visibility. The inner variable will take precedence within its scope, and any reference to the variable name in that scope will refer to the inner variable.
2. **Lifetime**: The outer variable still exists, and its value remains unchanged. Once you exit the inner scope, the outer variable becomes accessible again.
3. **Potential Confusion**: Shadowing can lead to confusion or bugs if a developer inadvertently modifies the inner variable and assumes it affects the outer one. This is especially important in larger codebases or nested functions.

## Why avoid variable shadowing?

1. **Readability Matters**: When you reuse variable names, it can get tricky to keep track of which variable you’re dealing with, especially in complex or nested code. This can make your code harder to read and understand at a glance.
2. **Maintenance Headaches**: If you come back to your code weeks or months later, shadowed variables can create confusion. You might forget which variable you were actually using or modifying, leading to potential bugs during updates or fixes.
3. **Debugging Woes**: If something goes wrong because of variable shadowing, it can take you longer to figure out what happened. You might not realize that you’re looking at the wrong variable, which adds frustration to the debugging process.
4. **Logical Slip-Ups**: It’s easy to mistakenly think you’re working with an outer variable when you’re actually using the inner one. This can lead to unexpected behavior in your code and might take a while to catch.

## How to avoid shadowing

1. **Clear Variable Names**: When naming your variables, go for descriptive names that reflect their purpose. This reduces the chances that you’ll need to reuse names, making your code clearer.
2. **Keep It Narrow**: Try to limit the scope of your variables as much as possible. The more specific you are about where a variable is declared, the less likely you are to accidentally shadow it.
3. **Be Careful with using :=**: In Go, using the := shorthand can easily lead to shadowing, especially in nested blocks. Be mindful of when you use it, and consider sticking with the = operator to assign values to existing variables instead.
4. **Use Linters**: Tools like linters can be incredibly helpful. They can flag shadowed variables and provide warnings, helping you maintain a high quality in your code.