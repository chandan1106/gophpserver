<?php
// Example API endpoint using GoPHPServer

// Set content type to JSON
header('Content-Type: application/json');

// Get request method
$method = $_SERVER['REQUEST_METHOD'];

// Sample data
$data = [
    'users' => [
        ['id' => 1, 'name' => 'John Doe', 'email' => 'john@example.com'],
        ['id' => 2, 'name' => 'Jane Smith', 'email' => 'jane@example.com'],
        ['id' => 3, 'name' => 'Bob Johnson', 'email' => 'bob@example.com'],
    ]
];

// Handle different HTTP methods
switch ($method) {
    case 'GET':
        // Return all users or a specific user
        $userId = isset($_GET['id']) ? (int)$_GET['id'] : null;
        
        if ($userId) {
            $user = null;
            foreach ($data['users'] as $u) {
                if ($u['id'] === $userId) {
                    $user = $u;
                    break;
                }
            }
            
            if ($user) {
                echo json_encode(['status' => 'success', 'data' => $user]);
            } else {
                http_response_code(404);
                echo json_encode(['status' => 'error', 'message' => 'User not found']);
            }
        } else {
            echo json_encode(['status' => 'success', 'data' => $data['users']]);
        }
        break;
        
    case 'POST':
        // Simulate creating a new user
        echo json_encode([
            'status' => 'success', 
            'message' => 'User created successfully',
            'data' => ['id' => 4, 'name' => 'New User', 'email' => 'new@example.com']
        ]);
        break;
        
    default:
        http_response_code(405);
        echo json_encode(['status' => 'error', 'message' => 'Method not allowed']);
        break;
}
?>