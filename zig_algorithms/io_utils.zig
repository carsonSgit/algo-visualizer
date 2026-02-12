const std = @import("std");

pub const Input = struct {
    array: []i64,
    algorithm: []const u8 = "",
};

pub const Step = struct {
    array: []const i64,
    comparing: []const usize,
    swapped: []const usize,
    sorted: []const usize,
    step_number: usize,
    message: []const u8,
};

pub fn readInput(allocator: std.mem.Allocator) !std.json.Parsed(Input) {
    const stdin = std.io.getStdIn().reader();
    const input_bytes = try stdin.readAllAlloc(allocator, 1024 * 1024); // 1MB limit
    defer allocator.free(input_bytes);
    return std.json.parseFromSlice(Input, allocator, input_bytes, .{ .ignore_unknown_fields = true });
}

pub fn printStep(
    array: []const i64,
    comparing: []const usize,
    swapped: []const usize,
    sorted: []const usize,
    step_number: usize,
    message: []const u8
) !void {
    const step = Step{
        .array = array,
        .comparing = comparing,
        .swapped = swapped,
        .sorted = sorted,
        .step_number = step_number,
        .message = message,
    };
    
    const stdout = std.io.getStdOut().writer();
    try std.json.stringify(step, .{}, stdout);
    try stdout.writeByte('\n');
}
