<?php
header('Content-Type: text/html');
?>
<!DOCTYPE html>
<html>
<head>
    <title>GoPHPServer Test</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 40px;
            line-height: 1.6;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
            border: 1px solid #ddd;
            border-radius: 5px;
        }
        h1 {
            color: #333;
        }
        .info {
            background-color: #f8f9fa;
            padding: 15px;
            border-radius: 4px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>GoPHPServer is working!</h1>
        
        <div class="info">
            <h2>PHP Information</h2>
            <p>PHP Version: <?php echo phpversion(); ?></p>
            <p>Server Time: <?php echo date('Y-m-d H:i:s'); ?></p>
            <p>Server Software: <?php echo $_SERVER['SERVER_SOFTWARE'] ?? 'GoPHPServer'; ?></p>
        </div>
        
        <h2>Request Information</h2>
        <ul>
            <li>Request Method: <?php echo $_SERVER['REQUEST_METHOD']; ?></li>
            <li>Request URI: <?php echo $_SERVER['REQUEST_URI']; ?></li>
            <li>Remote Address: <?php echo $_SERVER['REMOTE_ADDR']; ?></li>
        </ul>
    </div>
</body>
</html>